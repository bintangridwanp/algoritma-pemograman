import java.util.Scanner;

public class permukaanBola {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Luas Permukaan Bola");

        System.out.println("Masukkan jari-jari bola: ");
        double jariJari = input.nextDouble();

        // Validasi input
        if (jariJari <= 0) {
            System.out.println("Jari-jari bola harus lebih besar dari 0.");
        } else {
            // Hitung luas permukaan bola
            double pi = 3.14159; // Nilai pi
            double luasPermukaan = 4 * pi * (jariJari * jariJari);

            System.out.printf("Luas Permukaan Bola dengan jari-jari %.2f adalah: %.2f\n", jariJari, luasPermukaan);
        }

    }
}
