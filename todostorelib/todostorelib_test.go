package todostorelib

import (
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func Test_TodoStore(t *testing.T) {
	storeName := "TESTNAME"
	testStore := NewTodoStore(storeName)
	t.Log("Given the need to test the creation of the new ToDo store i.e 'vault'")
	{
		if testStore.name != storeName {
			t.Errorf("\t%s\tThe store name correct: Got: %q, Want: %q", failed, testStore.name, storeName)
		} else {
			t.Logf("\t%s\tThe store name correct: Got: %q, Want: %q", succeed, testStore.name, storeName)
		}
	}
	t.Log("Given the need to test the creaction of the new ToDo note in the store")
	message1 := "message1 for testing"
	testStore.AddTodo(message1)
	{
		response1, err := testStore.GetTodo(1)
		if err != nil {
			t.Fatalf("\t%s\tThe Todo note wasn't created", failed)
		} else {
			t.Logf("\t%s\tThe Todo note was created successfully", succeed)
		}
		if response1["message"] != message1 {
			t.Errorf("\t%s\tThe Todo note messages don't match", failed)
			t.Errorf("\t%s\tWant: %q", failed, message1)
			t.Errorf("\t%s\tGot: %q", failed, response1["message"])
		} else {
			t.Logf("\t%s\tThe Todo note messages match", succeed)
			t.Logf("\t%s\tWant: %q", succeed, message1)
			t.Logf("\t%s\tGot: %q", succeed, response1["message"])
		}
	}
	t.Log("Given a need to test the method returning vault info")
	{
		vaultInfo := testStore.GetVaultInfo()
		if vaultInfo["counter"] != "1" {
			t.Errorf("\t%s\tThe counter isn't showing the correct count", failed)
			t.Errorf("\t%s\tWant: %q", failed, "1")
			t.Errorf("\t%s\tGot: %q", failed, vaultInfo["counter"])
		} else {
			t.Logf("\t%s\tThe counter is showing the correct count", succeed)
			t.Logf("\t%s\tWant: %q", succeed, "1")
			t.Logf("\t%s\tGot: %q", succeed, vaultInfo["counter"])
		}

	}

}
