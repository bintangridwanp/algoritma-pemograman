import java.util.Scanner;

public class menentukanBilanganTerbesar {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menentukan Bilangan Terbesar");

        System.out.print("Masukkan bilangan pertama: ");
        double bilanganPertama = input.nextDouble();

        System.out.print("Masukkan bilangan kedua: ");
        double bilanganKedua = input.nextDouble();

        System.out.print("Masukkan bilangan ketiga: ");
        double bilanganKetiga = input.nextDouble();

        // Menentukan bilangan terbesar
        if (bilanganPertama > bilanganKetiga && bilanganPertama > bilanganKedua) {
            System.out.println("Bilangan terbesar adalah: " + bilanganPertama);
        } else if (bilanganKedua > bilanganPertama && bilanganKedua > bilanganKetiga) {
            System.out.println("Bilangan terbesar adalah: " + bilanganKedua);
        } else if (bilanganKetiga > bilanganPertama && bilanganKetiga > bilanganKedua) {
            System.out.println("Bilangan terbesar adalah: " + bilanganKetiga);
        } else {
            System.out.println("Ada dua atau lebih bilangan yang sama besar.");
        }
    }
}
