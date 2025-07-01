public class penjadwalanMeeting {
    public static void main(String[] args) {
        // Array 2 dimensi: [mulai, selesai]
        int[][] meetings = {
                {1, 4},
                {3, 5},
                {0, 6},
                {5, 7},
                {8, 9},
                {5, 9}
        };

        // Urutkan meeting berdasarkan waktu selesai (ascending) - bubble sort
        for (int i = 1; i < meetings.length; i++) {
            for (int j = 1; j < meetings.length - i; j++) {
                if (meetings[j][1] > meetings[j + 1][1]) {
                    // Tukar meeting[j] dan meeting[j+1]
                    int tempMulai = meetings[j][0];
                    int tempSelesai = meetings[j][1];

                    meetings[j][0] = meetings[j + 1][0];
                    meetings[j][1] = meetings[j + 1][1];

                    meetings[j + 1][0] = tempMulai;
                    meetings[j + 1][1] = tempSelesai;
                }
            }
        }

        // Gunakan greedy: pilih meeting yang tidak tumpang tindih
        int count = 1;
        int waktuBerakhirTerakhir = meetings[0][1];

        for (int i = 1; i < meetings.length; i++) {
            if (meetings[i][0] >= waktuBerakhirTerakhir) {
                count++;
                waktuBerakhirTerakhir = meetings[i][1];
            }
        }

        System.out.println("Maksimum meeting yang dapat dihadiri: " + count);
    }
}
