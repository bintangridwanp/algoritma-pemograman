import java.util.Scanner;

public class menukarNilaiVariable {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.println("Program Menukar Nilai Dua Variabel");

        System.out.print("Masukkan nilai variabel a: ");
        int a = input.nextInt();
        System.out.println("Nilai variabel a adalah: " + a + ", Berhasil disimpan.");

        System.out.print("Masukkan nilai variabel b: ");
        int b = input.nextInt();
        System.out.println("Nilai variabel b adalah: " + b + ", Berhasil disimpan.");

        // Menukar nilai variabel a dan b
        int temp = a;
        a = b;
        b = temp;
        System.out.println("Setelah penukaran:");
        System.out.println("Nilai variabel a sekarang adalah: " + a);
        System.out.println("Nilai variabel b sekarang adalah: " + b);

    }
}
