package aputil
import (
	"fmt"
	_"strconv"
	_"time"
	_"log"
	"os/exec"
	"os"
	"bufio"
	"strings"
	_"math/rand"
	_"time"
	"io/ioutil"
	
)


func ExecuteCommand_Output(cmdName string, cmdArgs []string) ([]byte, error){
	return exec.Command(cmdName, cmdArgs...).CombinedOutput()
	// if err != nil {
	// 	return out, err
	// }
	// fmt.Printf("utilllllsllllll%s  | %s\n",strings.ToUpper(cmdName), out)
	// return out, err
}


// func ExecuteCommand_withoutput(cmdName string, cmdArgs []string) error{
// 	cmd := exec.Command(cmdName, cmdArgs...)
// 	cmd.Stdin = strings.NewReader("some input")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("in all caps: %q\n", out.String())
// }



func ExecuteCommand(cmdName string, cmdArgs []string) error{
	//start := time.Now()
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s  | %s\n",strings.ToUpper(cmdName), scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return err
	}

	return nil

	// elapsed := time.Since(start)
	// fmt.Printf("%s Time: %s", cmdName, elapsed)

}

func WriteFile(filePath string, data []byte, mode os.FileMode) error {
	err := ioutil.WriteFile(filePath, data, mode)
	return err
}


