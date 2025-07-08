import java.util.Scanner;

public class konversiJamMenit {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Konversi Menit ke Jam");

        System.out.print("Masukkan total menit: ");
        int totalMenit = input.nextInt();

        // Menghitung jam dan menit
        int jam = totalMenit / 60;
        int sisaMenit = totalMenit % 60;

        // Menampilkan hasil konversi
        System.out.print("Total " + totalMenit + " menit adalah " + jam + " jam dan " + sisaMenit + " menit.");



    }
}
