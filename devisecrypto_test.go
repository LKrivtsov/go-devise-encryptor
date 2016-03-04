package devisecrypto

import (
  "testing"
)

func TestDigest(t *testing.T) {
  t.Log("Testing digest with a pepper")

  password,stretches,pepper := "changeme",4,"a bad pepper"

  hashedPassword,_ := Digest(password, stretches, pepper)
  anotherHashedPassword,_ := Digest(password, stretches, pepper)

  if hashedPassword == anotherHashedPassword {
    t.Errorf("The hashedPasswords should not be equal given two identical inputs. First password [%s] was equal to the second [%s] when a pepper is not present",hashedPassword,anotherHashedPassword)
  }

  pepper = ""
  t.Log("Testing digest with no pepper")

  hashedPassword,_ = Digest(password, stretches, pepper)
  anotherHashedPassword,_ = Digest(password, stretches, pepper)

  if hashedPassword == anotherHashedPassword {
    t.Errorf("The hashedPasswords should not be equal given two identical inputs. First password [%s] was equal to the second [%s] when a pepper is not present",hashedPassword,anotherHashedPassword)
  }

}

// Compare  compares a password with a salt against a hashed password.
func TestCompare(t *testing.T)  {
  t.Log("Testing compare with a pepper")

  password,stretches,pepper := "changeme",4,"a bad pepper"

  hashedPassword,_ := Digest(password, stretches, pepper)

  if Compare(password, pepper, hashedPassword) == false {
    t.Errorf("The hashed password should be identical to the original when given the same inputs and internal salt")
  }

  pepper = ""

  hashedPassword,_ = Digest(password, stretches, pepper)

  t.Log("Testing compare with no pepper")

  if Compare(password, pepper, hashedPassword) == false {
    t.Errorf("The hashed password should be identical to the original when given the same inputs and internal salt, when a pepper is not present")
  }

  t.Log("Testing compare with no hashedPassword provided")

  if Compare(password, pepper, "") != false {
    t.Errorf("If the hashedPassword is blank it should return false")
  }

}
