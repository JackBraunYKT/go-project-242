package code

import "testing"

func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		recursive bool
		human     bool
		all       bool
		want      string
		wantErr   bool
	}{
		{
			name:      "Test for test_data_1.csv",
			filePath:  "testdata/test_data_1.csv",
			recursive: false,
			human:     false,
			all:       false,
			want:      "947B",
			wantErr:   false,
		},
		{
			name:      "Test for test_data_2.csv",
			filePath:  "testdata/test_data_2.csv",
			recursive: false,
			all:       false,
			human:     false,
			want:      "1309B",
			wantErr:   false,
		},
		{
			name:      "Test for existing directory",
			filePath:  "testdata",
			recursive: false,
			all:       false,
			human:     false,
			want:      "2256B",
			wantErr:   false,
		},
		{
			name:      "Test for non-existing file",
			filePath:  "testdata/nonexistent.txt",
			recursive: false,
			all:       false,
			human:     false,
			want:      "0B",
			wantErr:   false,
		},
		{
			name:      "Test for inner directory and human-readable format",
			filePath:  "testdata/test-dir",
			recursive: false,
			all:       false,
			human:     true,
			want:      "2.5KB",
			wantErr:   false,
		},

		// All flag tests
		// hidden file
		{
			name:      "Test for hidden file with all flag",
			filePath:  "testdata/.hidden_test_data_3.csv",
			recursive: false,
			all:       true,
			human:     false,
			want:      "1309B",
			wantErr:   false,
		},
		{
			name:      "Test for hidden file with all flag and human-readable format",
			filePath:  "testdata/.hidden_test_data_3.csv",
			recursive: false,
			all:       true,
			human:     true,
			want:      "1.3KB",
			wantErr:   false,
		},

		// hidden file in directory
		{
			name:      "Test for directory with only hidden file without all flag",
			filePath:  "testdata/dir-with-hidden-file",
			recursive: false,
			all:       false,
			human:     false,
			want:      "0B",
			wantErr:   false,
		},
		{
			name:      "Test for directory with only hidden file with all flag",
			filePath:  "testdata/dir-with-hidden-file",
			recursive: false,
			all:       true,
			human:     false,
			want:      "1309B",
			wantErr:   false,
		},
		{
			name:      "Test for directory with only hidden file with all flag and human-readable format",
			filePath:  "testdata/dir-with-hidden-file",
			recursive: false,
			all:       true,
			human:     true,
			want:      "1.3KB",
			wantErr:   false,
		},

		{
			name:      "Test for hidden in directory with all flag and human-readable format",
			filePath:  "testdata/test-dir",
			recursive: false,
			all:       true,
			human:     true,
			want:      "2.5KB",
			wantErr:   false,
		},

		// Recursive flag tests
		{
			name:      "Test for recursive directory without all flag",
			filePath:  "testdata/recursive-dir",
			recursive: true,
			all:       false,
			human:     false,
			want:      "2586B",
			wantErr:   false,
		},
		{
			name:      "Test for recursive directory with all flag",
			filePath:  "testdata/recursive-dir",
			recursive: true,
			all:       true,
			human:     false,
			want:      "10043B",
			wantErr:   false,
		},
		{
			name:      "Test for recursive directory with all flag and human-readable format",
			filePath:  "testdata/recursive-dir",
			recursive: true,
			all:       true,
			human:     true,
			want:      "9.8KB",
			wantErr:   false,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			got, err := GetPathSize(testItem.filePath, testItem.recursive, testItem.human, testItem.all)
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
			want:  "512B",
		},
		{
			name:  "Test for kilobytes",
			bytes: 2048,
			want:  "2.0KB",
		},
		{
			name:  "Test for another kilobytes",
			bytes: 2458,
			want:  "2.4KB",
		},
		{
			name:  "Test for megabytes",
			bytes: 5 * 1024 * 1024,
			want:  "5.0MB",
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
		name      string
		filePath  string
		want      int64
		wantErr   bool
		all       bool
		recursive bool
	}{
		{
			name:      "Test for test_data_1.csv",
			filePath:  "testdata/test_data_1.csv",
			want:      947,
			wantErr:   false,
			recursive: false,
			all:       false,
		},
		{
			name:      "Test for test_data_2.csv",
			filePath:  "testdata/test_data_2.csv",
			want:      1309,
			wantErr:   false,
			recursive: false,
			all:       false,
		},
		{
			name:      "Test for existing directory",
			filePath:  "testdata",
			want:      2256,
			wantErr:   false,
			all:       false,
			recursive: false,
		},
		{
			name:      "Test for non-existing file",
			filePath:  "testdata/nonexistent.txt",
			want:      0,
			wantErr:   false,
			all:       false,
			recursive: false,
		},
		{
			name:      "Test for hidden file with all flag",
			filePath:  "testdata/.hidden_test_data_3.csv",
			want:      1309,
			wantErr:   false,
			all:       true,
			recursive: false,
		},
		{
			name:      "Test for recursive directory without all flag",
			filePath:  "testdata/recursive-dir",
			want:      2586,
			wantErr:   false,
			all:       false,
			recursive: true,
		},
		{
			name:      "Test for recursive directory with all flag",
			filePath:  "testdata/recursive-dir",
			want:      10043,
			wantErr:   false,
			all:       true,
			recursive: true,
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			got, err := GetSize(testItem.filePath, testItem.recursive, testItem.all)
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
