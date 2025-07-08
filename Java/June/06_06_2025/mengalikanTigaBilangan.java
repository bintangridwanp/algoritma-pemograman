import java.util.Scanner;

public class mengalikanTigaBilangan {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Mengalikan Tiga Bilangan");

        System.out.print("Masukkan bilangan pertama: ");
        double bilanganPertama = input.nextDouble();

        System.out.print("Masukkan bilangan kedua: ");
        double bilanganKedua = input.nextDouble();

        System.out.print("Masukkan bilangan ketiga: ");
        double bilanganKetiga = input.nextDouble();

        // Mengalikan ketiga bilangan
        double hasilPerkalian = bilanganPertama * bilanganKedua * bilanganKetiga;
        System.out.println("Hasil perkalian dari " + bilanganPertama + ", " + bilanganKedua + ", dan " + bilanganKetiga + " adalah " + hasilPerkalian + ".");

    }
}
