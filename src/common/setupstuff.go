package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	Dbsctpd      *sql.DB
	Dbloinc      *sql.DB
	Siteroot     string
	Spk          Spellchecker
	WordlistFile string
	// templates *Template
)

func check(e error, exit bool) {
	if e != nil {
		fmt.Println(e)
		if exit == true {
			panic(e)
			os.Exit(1)
		}
	}
}

func Initvars() {
	Siteroot = "..\\site\\"
	WordlistFile = "wordlist.txt"
}

func Setupdb() {
	var err error
	//SCT PD Database
	Dbsctpd, err = sql.Open("mysql", "root:@tcp(localhost:3307)/sct_pd")
	check(err, true)
	Dbsctpd.SetMaxIdleConns(100)
	// Open doesn't open a connection. Validate DSN data:
	err = Dbsctpd.Ping()
	check(err, true)

	//LOINC Database
	Dbloinc, err = sql.Open("mysql", "root:@tcp(localhost:3307)/loinc")
	check(err, true)
	Dbloinc.SetMaxIdleConns(100)
	// Open doesn't open a connection. Validate DSN data:
	err = Dbloinc.Ping()
	check(err, true)

}

func Setupspellchecker() {
	// TODO: Error if it dosent work - shut down
	Spk.Start(WordlistFile)
}

func Shutdown() {
	fmt.Println("shutting down...")
	Dbsctpd.Close()
	Dbloinc.Close()
	Spk.Close()
	os.Exit(0)
	fmt.Println("done, goodbye!")
}

/*
func SetupShutdown() {
	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		// Dbsctpd.Close()
		// Spk.Close()
		done <- true
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
}
*/
