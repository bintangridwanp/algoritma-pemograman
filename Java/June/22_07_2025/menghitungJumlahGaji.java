import java.util.Scanner;

public class menghitungJumlahGaji {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Jumlah Gaji Karyawan");

        System.out.println("Masukkan jumlah karyawan: ");
        int jumlahKaryawan = input.nextInt();
        double totalGaji = 0;

        for (int i = 1; i <= jumlahKaryawan; i++) {
            System.out.println("Karyawan ke-" + i);

            System.out.print("Masukkan gaji pokok karyawan: ");
            double gajiPokok = input.nextDouble();
            if (gajiPokok > 0) {
                totalGaji += gajiPokok;
            }
        }

        double rata2 = totalGaji / jumlahKaryawan;
        System.out.println("Jumlah gaji karyawan adalah: " + totalGaji);
        System.out.printf("Rata-rata gaji karyawan adalah: %.2f%n", rata2);

        input.close();
    }
}
