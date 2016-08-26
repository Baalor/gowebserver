package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    . "testing"
)

func TestRomanOne(t *T) {
    // We first create the http.Handler we wish to test
    n := romanGenerator(1)

    // We create an http.Request object to test with. The http.Request is
    // totally customizable in every way that a real-life http request is, so
    // even the most intricate behavior can be tested
    r, _ := http.NewRequest("GET", "/roman/1", nil)

    // httptest.Recorder implements the http.ResponseWriter interface, and as
    // such can be passed into ServeHTTP to receive the response. It will act as
    // if all data being given to it is being sent to a real client, when in
    // reality it's being buffered for later observation
    w := httptest.NewRecorder()

    // Pass in our httptest.Recorder and http.Request to our numberDumper. At
    // this point the numberDumper will act just as if it was responding to a
    // real request
    n.ServeHTTP(w, r)

    // httptest.Recorder gives a number of fields and methods which can be used
    // to observe the response made to our request. Here we check the response
    // code
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }

    // We can also get the full body out of the httptest.Recorder, and check
    // that its contents are what we expect
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: I\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}


func TestRomanTwo(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/2", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: II\n") {
        t.Fatalf("wrong body returned: %s", body)
    }
}

func TestRoman3(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/3", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: III\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman4(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/4", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: IV\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman5(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/5", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: V\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman6(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/6", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: VI\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman21(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/21", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: XXI\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}
func TestRoman50(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/50", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: L\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman500(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/500", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: D\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman2000(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/2000", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: MM\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman1990(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/1990", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: MCMXC\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}
func TestRoman2008(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/2008", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: MMVIII\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman99(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/99", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: XCIX\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman47(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/47", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: XLVII\n") {
        t.Fatalf("wrong body returned: %s", body)
    }

}

func TestRoman47unit(t *T) {
    n := to_roman(1)
    if n != "I" {
        t.Fatalf("wrong code returned: %d", 0)
    }
}


func TestRoman0(t *T) {
    n := romanGenerator(1)
    r, _ := http.NewRequest("GET", "/roman/0", nil)
    w := httptest.NewRecorder()
    n.ServeHTTP(w, r)
    if w.Code != 200 {
        t.Fatalf("wrong code returned: %d", w.Code)
    }
    body := w.Body.String()
    if body != fmt.Sprintf("Here's your number: NIL\n") {
        t.Fatalf("wrong body returned: %s", body)
    }
}