package main
import ("testing";"os";"net/http")

func TestPkgsrc(t *testing.T) {
	os.Setenv("CC","gcc")
	os.Setenv("MAKE_JOBS","4")
	os.Setenv("PKGSRC_TAR_GZ","http://localhost:8080/trunk.tar.gz")

	go func(){
		fs := http.FileServer(http.Dir("."))
		http.Handle("/",fs)
		http.ListenAndServe(":8080", nil)
	}()
	
	if err := pkgsrc(); err != nil {
		t.Error(err)
	}
}

func TestShowDepends(t *testing.T) {
	if err := run("show-depends"); err != nil {
		t.Error(err)
	}
}

func TestShowOptions(t *testing.T) {
	if err := run("show-options"); err != nil {
		t.Error(err)
	}
}

func TestInstall(t *testing.T) {
	if err := run("install","benchmarks/fib"); err != nil {
		t.Error(err)
	}		
}

func TestReplace(t *testing.T) {
	if err := run("replace","benchmarks/fib"); err != nil {
		t.Error(err)
	}	
}

func TestClean(t *testing.T) {	
	if err := run("clean","benchmarks/fib"); err != nil {
		t.Error(err)
	}
}

func TestCleanDepends(t *testing.T) {
	if err := run("clean-depends","benchmarks/fib"); err != nil {
		t.Error(err)
	}
}

func TestDeinstall(t *testing.T) {
	if err := run("deinstall","benchmarks/fib"); err != nil {
		t.Error(err)
	}
}