package code

import "testing"

func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		human    bool
		want     string
		wantErr  bool
	}{
		{
			name:     "Test for test_data_1.csv",
			filePath: "testdata/test_data_1.csv",
			human:    false,
			want:     "947\ttestdata/test_data_1.csv",
			wantErr:  false,
		},
		{
			name:     "Test for test_data_2.csv",
			filePath: "testdata/test_data_2.csv",
			human:    false,
			want:     "1309\ttestdata/test_data_2.csv",
			wantErr:  false,
		},
		{
			name:     "Test for existing directory",
			filePath: "testdata",
			human:    false,
			want:     "2256\ttestdata",
			wantErr:  false,
		},
		{
			name:     "Test for inner directory",
			filePath: "testdata/test-dir",
			human:    true,
			want:     "2.53 KB\ttestdata/test-dir",
			wantErr:  false,
		},
		{
			name:     "Test for non-existing file",
			filePath: "testdata/nonexistent.txt",
			human:    false,
			want:     "ошибка при чтении файла или директории",
			wantErr:  true,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			got, err := GetPathSize(testItem.filePath, testItem.human)
			if (err != nil) != testItem.wantErr {
				t.Errorf("GetSize() error = %v, wantErr %v", err, testItem.wantErr)
				return
			}
			if got != testItem.want {
				t.Errorf("GetSize() = %v, want %v", got, testItem.want)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name  string
		bytes float64
		want  string
	}{
		{
			name:  "Test for bytes",
			bytes: 512,
			want:  "512.00 B",
		},
		{
			name:  "Test for kilobytes",
			bytes: 2048,
			want:  "2.00 KB",
		},
		{
			name:  "Test for megabytes",
			bytes: 5 * 1024 * 1024,
			want:  "5.00 MB",
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			got := FormatSize(testItem.bytes)
			if got != testItem.want {
				t.Errorf("FormatSize() = %v, want %v", got, testItem.want)
			}
		})
	}
}

func TestGetSize(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     int64
		wantErr  bool
	}{
		{
			name:     "Test for test_data_1.csv",
			filePath: "testdata/test_data_1.csv",
			want:     947,
			wantErr:  false,
		},
		{
			name:     "Test for test_data_2.csv",
			filePath: "testdata/test_data_2.csv",
			want:     1309,
			wantErr:  false,
		},
		{
			name:     "Test for existing directory",
			filePath: "testdata",
			want:     2256,
			wantErr:  false,
		},
		{
			name:     "Test for non-existing file",
			filePath: "testdata/nonexistent.txt",
			want:     0,
			wantErr:  true,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			got, err := GetSize(testItem.filePath)
			if (err != nil) != testItem.wantErr {
				t.Errorf("GetSize() error = %v, wantErr %v", err, testItem.wantErr)
				return
			}
			if got != testItem.want {
				t.Errorf("GetSize() = %v, want %v", got, testItem.want)
			}
		})
	}
}
