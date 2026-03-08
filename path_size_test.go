package code

import "testing"

func TestGetSize(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     string
		wantErr  bool
	}{
		{
			name:     "Test for test_data_1.csv",
			filePath: "testdata/test_data_1.csv",
			want:     "947B\ttestdata/test_data_1.csv",
			wantErr:  false,
		},
		{
			name:     "Test for test_data_2.csv",
			filePath: "testdata/test_data_2.csv",
			want:     "1309B\ttestdata/test_data_2.csv",
			wantErr:  false,
		},
		{
			name:	 "Test for existing directory",
			filePath: "testdata",
			want:     "2256B\ttestdata",
			wantErr:  false,
		},
		{
			name:    "Test for inner directory",
			filePath: "testdata/test-dir",
			want:     "2586B\ttestdata/test-dir",
			wantErr:  false,
		},
		{
			name:     "Test for non-existing file",
			filePath: "testdata/nonexistent.txt",
			want:     "",
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