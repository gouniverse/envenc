package envenc

import "testing"

func TestObfuscate(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		// want string
	}{
		{
			name: "Hello world",
			args: args{
				input: "Hello world",
			},
		},
		{
			name: "Lorem ipsum",
			args: args{
				input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus nunc mauris, hendrerit vitae erat nec, aliquam vehicula nulla. Praesent turpis risus, varius nec auctor vitae, consequat quis nisl. Donec nisl diam, dictum id arcu at, pellentesque consequat elit. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Sed a cursus tortor. Sed eget sagittis velit. Praesent convallis ultrices neque, eget aliquam lectus varius ut.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1, err := Obfuscate(tt.args.input)

			if err != nil {
				t.Errorf("Obfuscate() error = %v", err)
			}

			t.Log(str1)

			str2, err := Deobfuscate(str1)

			if err != nil {
				t.Errorf("Deobfuscate() error = %v", err)
			}

			if str2 != tt.args.input {
				t.Errorf("Deobfuscate() = %v, want %v", str2, tt.args.input)
			}
		})
	}
}
