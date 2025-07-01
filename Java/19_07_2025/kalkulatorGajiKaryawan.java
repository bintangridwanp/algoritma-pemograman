import java.util.Scanner;

public class kalkulatorGajiKaryawan {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Kalkulator Gaji Karyawan");

        System.out.print("Masukkan gaji pokok karyawan: ");
        double gajiPokok = input.nextDouble();

        System.out.print("Masukkan jumlah jam lembur karyawan: ");
        double jamLembur = input.nextDouble();

        System.out.print("Masukkan tarif lembur per jam: ");
        double tarifLembur = input.nextDouble();

        // Menghitung total gaji
        double totalGaji = gajiPokok + (jamLembur * tarifLembur);
        System.out.println("Total gaji karyawan adalah: " + totalGaji);

        
    }
}
