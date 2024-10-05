package main

import "testing"

func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}

  t.Run("search value", func(t *testing.T) {
    got, _ := dictionary.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
  })

  t.Run("search unknown value", func(t *testing.T) {
    _, err := dictionary.Search("unkown")
    want := ErrKeyNotFound 

    if err == nil {
      t.Fatal("expected to get an error")
    }

    assertStrings(t, err.Error(), want.Error())
  })
}

func TestAdd(t *testing.T) {
  t.Run("add new key", func(t *testing.T) {
    dictionary := Dictionary{}
    key := "test"
    value := "this is just a test"

    err := dictionary.Add(key, value) 

    assertErr(t, err, nil)
    assertValue(t, dictionary, key, value)
  })

  t.Run("add existing key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionary := Dictionary{key: value}

    err := dictionary.Add(key, "new test") 

    assertErr(t, err, ErrKeyExists)
    assertValue(t, dictionary, key, value)
  })
}

func TestUpdate(t *testing.T) {
  t.Run("update value", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionaty := Dictionary{key: value}
    newValue := "new value"

    err := dictionaty.Update(key, newValue)

    assertErr(t, err, nil)
    assertValue(t, dictionaty, key, newValue)
  })

  t.Run("update with key does not exist", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionaty := Dictionary{key: value}
    newValue := "new value"
    keyDoesNotExist := "keyDoesNotExist"

    err := dictionaty.Update(keyDoesNotExist, newValue)

    assertErr(t, err, ErrUpdateKeyDoesNotExist)
  })
}

func TestDelete(t *testing.T) {
  t.Run("delete key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionaty := Dictionary{key: value}

    dictionaty.Delete(key)

    _, err := dictionaty.Search(key)
    assertErr(t, err, ErrKeyNotFound)
  })

  t.Run("delete key that does not exist", func(t *testing.T) {
    key := "test"
    dictionaty := Dictionary{}

    err := dictionaty.Delete(key)

    assertErr(t, err, ErrDeleteKeyDoesNotExist)
  })
}

//=================== Helper functions ====================== 

func assertStrings(t testing.TB, got, want string) {
  t.Helper()

  if got != want {
    t.Errorf("got: %s, want: %s", got, want)
  }
}

func assertValue(t testing.TB, dictionary Dictionary, key, value string) {
  t.Helper()

  got, err := dictionary.Search(key)

  if err != nil {
    t.Fatal("should find added word: ", err)
  }
  assertStrings(t, got, value)
}

func assertErr(t testing.TB, got, want error) {
  t.Helper()

  if got != want {
    t.Errorf("got: %s, want: %s", got.Error(), want.Error())
  }
}
