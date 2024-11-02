package util

import (
	"log"
	"os/exec"
	"syscall"
)

// RunCommand runs a command and ensures that its child processes have the same lifecycle as the parent process.
// If the parent process exits, the child process will also be terminated.
func RunCommand(name string, args ...string) error {
	// Create a new process
	cmd := exec.Command(name, args...)

	// Set process group attributes
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, // Create a new process group
		Pgid:    0,    // The child process will become the leader of the new process group
	}

	// Start the subprocess
	if err := cmd.Start(); err != nil {
		return err
	}

	// Register an exit handler to clean up the child process when the parent process exits
	cleanup := func() {
		pgid, err := syscall.Getpgid(cmd.Process.Pid)
		if err != nil {
			log.Printf("Failed to get process group id: %v", err)
			return
		}
		if err := syscall.Kill(-pgid, syscall.SIGTERM); err != nil { // Send signal to the entire process group
			log.Printf("Failed to send signal to process group: %v", err)
		}
		if err := cmd.Wait(); err != nil {
			log.Printf("Command finished with error: %v", err)
		}
	}

	// Register the exit handler
	defer cleanup()

	// Wait for the subprocess to finish
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
