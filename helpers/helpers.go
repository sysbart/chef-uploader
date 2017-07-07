package helpers
import (
	"os/exec"
  "log"
)
// Based on https://www.dotnetperls.com/duplicates-go
func RemoveDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v:= range elements {
        encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range encountered {
        result = append(result, key)
    }
    return result
}


func RunCommand(cmd string, args...string) []byte {
  runCmd, err := exec.Command(cmd, args...).Output()
  if err != nil {
    log.Print(string(runCmd))
    log.Fatal(err)
  }
  return runCmd
}
