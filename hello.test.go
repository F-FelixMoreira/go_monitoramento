func TestPrintGreeting(t *testing.T) {
	var buf bytes.Buffer
	fmt.Println = func(a ...interface{}) (n int, err error) {
		return fmt.Fprint(&buf, a...)
	}

	main()

	greeting := "Olá,\n"
	if buf.String() != greeting {
		t.Errorf("Expected '%s', got '%s'", greeting, buf.String())
	}
}func TestIdade(t *testing.T) {
	var idade int = 24
	output := fmt.Sprintf("Sua idade é %d", idade)
	expected := "sua idade é 24"
	if output != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, output)
	}
}