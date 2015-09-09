package volume_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/baggageclaim/fs"
	"github.com/concourse/baggageclaim/volume"
	"github.com/concourse/baggageclaim/volume/driver"
	"github.com/concourse/baggageclaim/volume/fakes"
	"github.com/pivotal-golang/lager/lagertest"
)

var _ = Describe("Repository", func() {
	var (
		volumeDir  string
		fakeLocker *fakes.FakeLockManager
	)

	zero := uint(0)

	BeforeEach(func() {
		var err error
		volumeDir, err = ioutil.TempDir("", fmt.Sprintf("baggageclaim_volume_dir_%d", GinkgoParallelNode()))
		Ω(err).ShouldNot(HaveOccurred())

		fakeLocker = new(fakes.FakeLockManager)
	})

	AfterEach(func() {
		err := os.RemoveAll(volumeDir)
		Ω(err).ShouldNot(HaveOccurred())
	})

	Describe("naive", func() {
		var (
			fakeDriver *fakes.FakeDriver
			repo       volume.Repository

			someVolume volume.Volume
		)

		BeforeEach(func() {
			fakeDriver = new(fakes.FakeDriver)
			logger := lagertest.NewTestLogger("repo")
			repo = volume.NewRepository(logger, fakeDriver, fakeLocker, volumeDir, volume.TTL(60))

			var err error
			someVolume, err = repo.CreateVolume(volume.Strategy{
				"type": volume.StrategyEmpty,
			}, volume.Properties{}, &zero)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Describe("destroying a volume", func() {
			It("calls DestroyVolume on the driver", func() {
				Ω(filepath.Join(volumeDir, someVolume.Handle)).Should(BeADirectory())

				err := repo.DestroyVolume(someVolume.Handle)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeDriver.DestroyVolumeCallCount()).Should(Equal(1))
				volumePath := fakeDriver.DestroyVolumeArgsForCall(0)
				Ω(volumePath).Should(Equal(someVolume.Path))
			})

			It("deletes it from the disk", func() {
				volumes, err := repo.ListVolumes(volume.Properties{})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(volumes).Should(HaveLen(1))

				Ω(filepath.Join(volumeDir, someVolume.Handle)).Should(BeADirectory())

				err = repo.DestroyVolume(someVolume.Handle)
				Ω(err).ShouldNot(HaveOccurred())

				volumes, err = repo.ListVolumes(volume.Properties{})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(volumes).Should(HaveLen(0))

				Ω(filepath.Join(volumeDir, someVolume.Handle)).ShouldNot(BeADirectory())
			})

			It("removes it from listVolumes", func() {
				Ω(filepath.Join(volumeDir, someVolume.Handle)).Should(BeADirectory())

				err := repo.DestroyVolume(someVolume.Handle)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(repo.ListVolumes(volume.Properties{})).Should(BeEmpty())
			})

			It("makes some attempt at locking", func() {
				err := repo.DestroyVolume(someVolume.Handle)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeLocker.LockCallCount()).Should(Equal(1))
				Ω(fakeLocker.LockArgsForCall(0)).Should(Equal(someVolume.Handle))
				Ω(fakeLocker.UnlockCallCount()).Should(Equal(1))
				Ω(fakeLocker.UnlockArgsForCall(0)).Should(Equal(someVolume.Handle))
			})
		})

		Describe("setting properties on a volume", func() {
			It("makes some attempt at locking", func() {
				err := repo.SetProperty(someVolume.Handle, "hello", "goodbye")
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeLocker.LockCallCount()).Should(Equal(1))
				Ω(fakeLocker.LockArgsForCall(0)).Should(Equal(someVolume.Handle))
				Ω(fakeLocker.UnlockCallCount()).Should(Equal(1))
				Ω(fakeLocker.UnlockArgsForCall(0)).Should(Equal(someVolume.Handle))
			})
		})

		Describe("setting the TTL on a volume", func() {
			It("makes some attempt at locking", func() {
				err := repo.SetTTL(someVolume.Handle, 42)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeLocker.LockCallCount()).Should(Equal(1))
				Ω(fakeLocker.LockArgsForCall(0)).Should(Equal(someVolume.Handle))
				Ω(fakeLocker.UnlockCallCount()).Should(Equal(1))
				Ω(fakeLocker.UnlockArgsForCall(0)).Should(Equal(someVolume.Handle))
			})
		})
	})

	Describe("BtrFS", func() {
		if runtime.GOOS != "linux" {
			fmt.Println("\x1b[33m*** skipping btrfs tests because non-linux ***\x1b[0m")
			return
		}

		var (
			tempDir    string
			volumeDir  string
			fsDriver   *driver.BtrFSDriver
			filesystem *fs.BtrfsFilesystem
			repo       volume.Repository

			someVolume volume.Volume
		)

		BeforeEach(func() {
			var err error
			tempDir, err = ioutil.TempDir("", "baggageclaim_repo_test")
			Ω(err).ShouldNot(HaveOccurred())

			dLogger := lagertest.NewTestLogger("driver")
			fsDriver = driver.NewBtrFSDriver(dLogger)

			imagePath := filepath.Join(tempDir, "image.img")
			volumeDir = filepath.Join(tempDir, "mountpoint")
			filesystem = fs.New(dLogger, imagePath, volumeDir)
			err = filesystem.Create(100 * 1024 * 1024)
			Ω(err).ShouldNot(HaveOccurred())

			logger := lagertest.NewTestLogger("repo")
			repo = volume.NewRepository(logger, fsDriver, fakeLocker, volumeDir, volume.TTL(60))

			someVolume, err = repo.CreateVolume(volume.Strategy{
				"type": volume.StrategyEmpty,
			}, volume.Properties{}, &zero)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := filesystem.Delete()
			Ω(err).ShouldNot(HaveOccurred())

			err = os.RemoveAll(tempDir)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Describe("creating a new volume", func() {
			It("cows", func() {
				childVolume, err := repo.CreateVolume(volume.Strategy{
					"type":   volume.StrategyCopyOnWrite,
					"volume": someVolume.Handle,
				}, volume.Properties{}, &zero)
				Ω(err).ShouldNot(HaveOccurred())

				childFilePath := filepath.Join(childVolume.Path, "this-should-only-be-in-the-child")
				err = ioutil.WriteFile(childFilePath, []byte("contents"), 0755)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(childFilePath).Should(BeARegularFile())

				parentFilePath := filepath.Join(someVolume.Path, "this-should-only-be-in-the-child")
				Ω(parentFilePath).ShouldNot(BeADirectory())
			})

			It("makes some attempt at locking the parent", func() {
				_, err := repo.CreateVolume(volume.Strategy{
					"type":   volume.StrategyCopyOnWrite,
					"volume": someVolume.Handle,
				}, volume.Properties{}, &zero)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeLocker.LockCallCount()).Should(Equal(1))
				Ω(fakeLocker.LockArgsForCall(0)).Should(Equal(someVolume.Handle))
				Ω(fakeLocker.UnlockCallCount()).Should(Equal(1))
				Ω(fakeLocker.UnlockArgsForCall(0)).Should(Equal(someVolume.Handle))
			})
		})

		Describe("destroying a volume", func() {
			It("deletes it", func() {
				Ω(filepath.Join(volumeDir, someVolume.Handle)).Should(BeADirectory())

				err := repo.DestroyVolume(someVolume.Handle)
				Ω(err).ShouldNot(HaveOccurred())

				volumes, err := repo.ListVolumes(volume.Properties{})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(volumes).Should(HaveLen(0))

				Ω(filepath.Join(volumeDir, someVolume.Handle)).ShouldNot(BeADirectory())
			})

			It("makes some attempt at locking", func() {
				someVolume, err := repo.CreateVolume(volume.Strategy{
					"type": volume.StrategyEmpty,
				}, volume.Properties{}, &zero)
				Ω(err).ShouldNot(HaveOccurred())

				err = repo.DestroyVolume(someVolume.Handle)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeLocker.LockCallCount()).Should(Equal(1))
				Ω(fakeLocker.LockArgsForCall(0)).Should(Equal(someVolume.Handle))
				Ω(fakeLocker.UnlockCallCount()).Should(Equal(1))
				Ω(fakeLocker.UnlockArgsForCall(0)).Should(Equal(someVolume.Handle))
			})
		})

		Describe("setting properties on a volume", func() {
			It("makes some attempt at locking", func() {
				err := repo.SetProperty(someVolume.Handle, "hello", "goodbye")
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeLocker.LockCallCount()).Should(Equal(1))
				Ω(fakeLocker.LockArgsForCall(0)).Should(Equal(someVolume.Handle))
				Ω(fakeLocker.UnlockCallCount()).Should(Equal(1))
				Ω(fakeLocker.UnlockArgsForCall(0)).Should(Equal(someVolume.Handle))
			})
		})

		Describe("setting the TTL on a volume", func() {
			It("makes some attempt at locking", func() {
				err := repo.SetTTL(someVolume.Handle, 42)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeLocker.LockCallCount()).Should(Equal(1))
				Ω(fakeLocker.LockArgsForCall(0)).Should(Equal(someVolume.Handle))
				Ω(fakeLocker.UnlockCallCount()).Should(Equal(1))
				Ω(fakeLocker.UnlockArgsForCall(0)).Should(Equal(someVolume.Handle))
			})
		})
	})
})
