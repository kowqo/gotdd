package main

import "testing"

func TestSearch(t *testing.T) {
	assertValue := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Fatal("Don't wanna error but got one ", err)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("Want error but didn't get one", want)
		}

		if got != want {
			t.Errorf("want %s but got %s", got, want)
		}
	}

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Получить существующее", func(t *testing.T) {

		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertNoError(t, err)

		assertValue(t, got, want)
	})

	t.Run("Нет значения в мапе", func(t *testing.T) {

		_, err := dictionary.Search("no")

		assertError(t, err, UnknownWordError)
	})

	t.Run("Добавить слово", func(t *testing.T) {
		const novoe = "Novoe"
		const key = "test1"
		var err error
		var got string

		err = dictionary.Add(key, novoe)

		assertNoError(t, err)

		got, err = dictionary.Search(key)

		assertNoError(t, err)

		assertValue(t, got, novoe)
	})

	t.Run("Добавить слово ключ которой уже есть в словаре", func(t *testing.T) {
		const novoe = "Novoe"
		const key = "test1"
		var err error

		err = dictionary.Add(key, novoe)

		assertError(t, err, AlreadyHaveKeyValue)

	})

	t.Run("Обновляем значение слова", func(t *testing.T) {
		const key = "test1"
		const old = "Novoe"
		var err error
		var got string
		dict := Dictionary{key: old}
		newSlovo := "New slovo"

		err = dict.Update(key, newSlovo)
		assertNoError(t, err)

		got, err = dict.Search(key)
		assertNoError(t, err)

		assertValue(t, newSlovo, got)

	})

	t.Run("Обновляем не существуюещее словао", func(t *testing.T) {
		const key = "test1"
		dict := Dictionary{}
		newSlovo := "New slovo"
		var err error
		var got string

		err = dict.Update(key, newSlovo)
		assertError(t, err, DontHaveKey)

		got, err = dict.Search(key)

		assertError(t, err, UnknownWordError)

		assertValue(t, "", got)
	})

	t.Run("Удаляем слово которое есть в словаер", func(t *testing.T) {
		const key = "test1"
		dict := Dictionary{key: "slovo"}

		err := dict.Delete(key)

		assertNoError(t, err)
	})

	t.Run("Пытаемся удалить слово которого нет в словаре", func(t *testing.T) {
		dict := Dictionary{"Hello": "slovo"}

		err := dict.Delete("key")

		assertError(t, err, CantDeleteNonExistKey)
	})
}
