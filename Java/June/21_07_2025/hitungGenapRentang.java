import java.util.Scanner;

public class hitungGenapRentang {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Hitung Bilangan Genap dalam Rentang");

        System.out.print("Masukkan bilangan awal: ");
        int awal = input.nextInt();

        System.out.print("Masukkan bilangan akhir: ");
        int akhir = input.nextInt();

        // Validasi input
        if (awal > akhir) {
            System.out.println("Bilangan awal harus kurang dari atau sama dengan bilangan akhir.");
            return;
        }

        System.out.println("Bilangan genap dalam rentang " + awal + " sampai " + akhir + ":");
        int jumlahGenap = 0;
        for (int i = awal; i <= akhir; i++) {
            if (i % 2 == 0) {
                System.out.print(i + " ");
                jumlahGenap++;
            }
        }
        System.out.println("\nJumlah bilangan genap dalam rentang tersebut adalah: " + jumlahGenap);
        input.close();
    }
}
