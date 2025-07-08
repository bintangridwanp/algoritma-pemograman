import java.util.Scanner;

public class menghitungRataBeratKoper {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Rata-Rata Berat Koper");

        System.out.print("Masukkan jumlah koper yang ingin dihitung rata-ratanya: ");
        int jumlahKoper = input.nextInt();

        // Validasi input
        if (jumlahKoper <= 0) {
            System.out.println("Jumlah koper harus lebih dari 0.");
            return;
        }

        double totalBerat = 0.0;
        for (int i = 1; i <= jumlahKoper; i++) {
            System.out.print("Masukkan berat koper ke-" + i + " (dalam kg): ");
            double beratKoper = input.nextDouble();

            // Validasi berat koper
            if (beratKoper < 0) {
                System.out.println("Berat koper tidak boleh negatif.");
                return;
            }

            totalBerat += beratKoper;
        }

        double rataRataBerat = totalBerat / jumlahKoper;

        System.out.printf("Rata-rata berat koper adalah %.2f kg.%n", rataRataBerat);
        System.out.println("Terima kasih telah menggunakan program ini!");
    }
}
