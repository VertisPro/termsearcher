package common

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

type Spellchecker struct {
	cmd           *exec.Cmd
	inpipe        io.WriteCloser
	outpipe       io.ReadCloser
	errpipe       io.ReadCloser
	stdoutscanner *bufio.Scanner
	stderrscanner *bufio.Scanner
}

func (sp *Spellchecker) Start(WordlistFile string) error {
	var err error
	sp.cmd = exec.Command("python", "-i")
	sp.inpipe, err = sp.cmd.StdinPipe()
	check(err, true)
	sp.outpipe, err = sp.cmd.StdoutPipe()
	check(err, true)
	sp.errpipe, err = sp.cmd.StderrPipe()
	check(err, true)
	err = sp.cmd.Start()
	check(err, true)
	sp.stdoutscanner = bufio.NewScanner(sp.outpipe)
	sp.stderrscanner = bufio.NewScanner(sp.errpipe)
	_, err = sp.inpipe.Write([]byte("execfile('enchant_console.py')\r\n"))
	check(err, true)

	return nil
}

func (sp *Spellchecker) CheckText(sentence string) (errorwords string) {
	syntax := "print CheckText('" + sentence + "')"
	return string(sp.processcommand(syntax))
}

func (sp *Spellchecker) SuggestSpelling(word string) (suggestions string) {
	syntax := "print SuggestSpelling(\"" + word + "\")"
	return string(sp.processcommand(syntax))
}

func (sp *Spellchecker) processcommand(src string) []byte {
	var err error
	_, err = sp.inpipe.Write([]byte(src + "\r\n"))
	check(err, false)
	sp.stdoutscanner.Scan()
	sp.stderrscanner.Scan()
	if err = sp.stdoutscanner.Err(); err != nil {
		fmt.Println("reading standard out:", err)
	}
	if err = sp.stderrscanner.Err(); err != nil {
		fmt.Println("reading standard out:", err)
	}
	return sp.stdoutscanner.Bytes()
}

func (sp *Spellchecker) Close() {
	_ = sp.processcommand("exit()")
	sp.inpipe.Close()
	sp.outpipe.Close()
	sp.errpipe.Close()
	sp.cmd.Wait()
}

/*
func main() {
	binary := "python"
	args := []string{"-i"}
	wordlist := "wordlist.txt"
	setupspellchecker(binary, wordlist, args)
	defer cmd.Wait()
	fmt.Println(suggestpelling("diabet"))
	fmt.Println(suggestpelling("myocaedial"))
	fmt.Println(suggestpelling("fractire"))
	fmt.Println(suggestpelling("fracture"))
	fmt.Println(suggestpelling("artointestinal"))
	fmt.Println(suggestpelling("galctophritis"))

	// Wrapping the unbuffered `os.Stdin` with a buffered
	// scanner gives us a convenient `Scan` method that
	// advances the scanner to the next token; which is
	// the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// `Text` returns the current token, here the next line,
		// from the input.
		word := scanner.Text()
		if word == "exit()" {
			_ = processcommand("exit()")
			inpipe.Close()
			outpipe.Close()
			errpipe.Close()
			//cmd.Wait()
			os.Exit(0)
		} else {
			// Write out the uppercased line.
			fmt.Println(suggestpelling(word))
		}
	}
}
*/
