import java.util.Scanner;

public class mencariBolaPalingBerat {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Mencari Bola Paling Berat");

        System.out.print("Masukkan jumlah bola: ");
        int jumlahBola = input.nextInt();

        // Validasi input jumlah bola
        if(jumlahBola <= 0) {
            System.out.println("Jumlah bola haru lebih dari 0.");
            return;
        }

        int bolaTerberat = 0;
        for (int i = 1; i <= jumlahBola; i++) {
            System.out.print("Masukkan berat bola ke-" + i + ": ");
            int beratBola = input.nextInt();

            // Validasi berat bola
            if (beratBola < 0) {
                System.out.println("Berat bola tidak boleh negatif.");
                return;
            }

            // Mencari bola terberat
            if (beratBola > bolaTerberat) {
                bolaTerberat = beratBola;
            }
        }

        System.out.println("Bola terberat memiliki berat: " + bolaTerberat + " kg.");
        System.out.println("Terima kasih telah menggunakan program ini!");

    }
}
