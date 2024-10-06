package main

type Dictionary map[string]string

const (
  ErrKeyNotFound = DictionaryErr("could not find the key you were looking for")
  ErrKeyExists = DictionaryErr("cannot add, key already exists in the dictionary")
  ErrUpdateKeyDoesNotExist = DictionaryErr("cannot update, key does not exist in the dictionary")
  ErrDeleteKeyDoesNotExist = DictionaryErr("cannot update, key does not exist in the dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
  return string(e)
}

func (d Dictionary) Add(key, value string) error {
  _, err := d.Search(key)

  switch err {
  case ErrKeyNotFound:
    d[key] = value 
    return nil
  case nil:
    return ErrKeyExists
  default:
    return err
  }
} 

func (d Dictionary) Update(key, newValue string) error {
  _, err := d.Search(key)

  switch err {
  case ErrKeyNotFound:
    return ErrUpdateKeyDoesNotExist
  case nil:
    d[key] = newValue
    return nil
  default:
    return err
  }
} 

func (d Dictionary) Delete(key string) error {
  _, err := d.Search(key)

  switch err {
  case ErrKeyNotFound:
    return ErrDeleteKeyDoesNotExist
  case nil:
    delete(d, key)
    return nil
  default:
    return err
  }
}

func (d Dictionary) Search(key string) (string, error) {
  value, ok := d[key]

  if !ok {
    return "", ErrKeyNotFound
  }

  return value, nil
}

func main() {
}
