package code

import "testing"

func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		all      bool
		human    bool
		want     string
		wantErr  bool
	}{
		{
			name:     "Test for test_data_1.csv",
			filePath: "testdata/test_data_1.csv",
			human:    false,
			all:      false,
			want:     "947B\ttestdata/test_data_1.csv",
			wantErr:  false,
		},
		{
			name:     "Test for test_data_2.csv",
			filePath: "testdata/test_data_2.csv",
			all:      false,
			human:    false,
			want:     "1309B\ttestdata/test_data_2.csv",
			wantErr:  false,
		},
		{
			name:     "Test for existing directory",
			filePath: "testdata",
			all:      false,
			human:    false,
			want:     "2256B\ttestdata",
			wantErr:  false,
		},
		{
			name:     "Test for non-existing file",
			filePath: "testdata/nonexistent.txt",
			all:      false,
			human:    false,
			want:     "ошибка при чтении файла или директории",
			wantErr:  true,
		},
		{
			name:     "Test for inner directory and human-readable format",
			filePath: "testdata/test-dir",
			all:      false,
			human:    true,
			want:     "2.53KB\ttestdata/test-dir",
			wantErr:  false,
		},

		// All flag tests
		// hidden file
		{
			name:     "Test for hidden file with all flag",
			filePath: "testdata/.hidden_test_data_3.csv",
			all:      true,
			human:    false,
			want:     "1309B\ttestdata/.hidden_test_data_3.csv",
			wantErr:  false,
		},
		{
			name:     "Test for hidden file with all flag and human-readable format",
			filePath: "testdata/.hidden_test_data_3.csv",
			all:      true,
			human:    true,
			want:     "1.28KB\ttestdata/.hidden_test_data_3.csv",
			wantErr:  false,
		},

		// hidden file in directory
		{
			name:     "Test for directory with only hidden file without all flag",
			filePath: "testdata/dir-with-hidden-file",
			all:      false,
			human:    false,
			want:     "0B\ttestdata/dir-with-hidden-file",
			wantErr:  false,
		},
		{
			name:     "Test for directory with only hidden file with all flag",
			filePath: "testdata/dir-with-hidden-file",
			all:      true,
			human:    false,
			want:     "1309B\ttestdata/dir-with-hidden-file",
			wantErr:  false,
		},
		{
			name:     "Test for directory with only hidden file with all flag and human-readable format",
			filePath: "testdata/dir-with-hidden-file",
			all:      true,
			human:    true,
			want:     "1.28KB\ttestdata/dir-with-hidden-file",
			wantErr:  false,
		},

		{
			name:     "Test for hidden in directory with all flag and human-readable format",
			filePath: "testdata/test-dir",
			all:      true,
			human:    true,
			want:     "2.53KB\ttestdata/test-dir",
			wantErr:  false,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			got, err := GetPathSize(testItem.filePath, testItem.human, testItem.all)
			if (err != nil) != testItem.wantErr {
				t.Errorf("GetPathSize() error = %v, wantErr %v", err, testItem.wantErr)
				return
			}
			if got != testItem.want {
				t.Errorf("GetPathSize() = %v, want %v", got, testItem.want)
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
			want:  "512.00B",
		},
		{
			name:  "Test for kilobytes",
			bytes: 2048,
			want:  "2.00KB",
		},
		{
			name:  "Test for megabytes",
			bytes: 5 * 1024 * 1024,
			want:  "5.00MB",
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
		all      bool
		wantErr  bool
	}{
		{
			name:     "Test for test_data_1.csv",
			filePath: "testdata/test_data_1.csv",
			want:     947,
			all:      false,
			wantErr:  false,
		},
		{
			name:     "Test for test_data_2.csv",
			filePath: "testdata/test_data_2.csv",
			want:     1309,
			all:      false,
			wantErr:  false,
		},
		{
			name:     "Test for existing directory",
			filePath: "testdata",
			want:     2256,
			all:      false,
			wantErr:  false,
		},
		{
			name:     "Test for non-existing file",
			filePath: "testdata/nonexistent.txt",
			want:     0,
			all:      false,
			wantErr:  true,
		},
		{
			name:     "Test for hidden file with all flag",
			filePath: "testdata/.hidden_test_data_3.csv",
			want:     1309,
			all:      true,
			wantErr:  false,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			got, err := GetSize(testItem.filePath, testItem.all)
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
