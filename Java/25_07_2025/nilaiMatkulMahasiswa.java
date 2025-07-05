import java.util.Scanner;

public class nilaiMatkulMahasiswa {
    public static void main(String[] args) {
     Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Nilai Mata Kuliah Mahasiswa");

        System.out.print("Masukkan jumlah mahasiswa : ");
        int mahasiswa = input.nextInt();

        System.out.println("Masukkan mata kuliah yang diambil : ");
        int matkul = input.nextInt();

        for (int i = 0; i < mahasiswa; i++) {
            System.out.println("Mahasiswa ke-" + (i + 1) + " : ");
            double totalNilai = 0;

            for (int j = 0; j < matkul; j++) {
                System.out.print("Masukkan nilai mata kuliah ke-" + (j + 1) + " : ");
                double nilai = input.nextDouble();

                // Validasi nilai
                if (nilai < 0 || nilai > 100) {
                    System.out.println("Nilai harus antara 0 dan 100.");
                    j--;
                    continue;
                }
                totalNilai += nilai;
            }
            double rataRata = totalNilai / matkul;
            System.out.printf("Rata-rata nilai mahasiswa ke-%d adalah: %.2f%n", (i + 1), rataRata);
        }
    }
}
