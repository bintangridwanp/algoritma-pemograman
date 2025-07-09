import java.util.Scanner;

public class menentukanBilanganUnik {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menentukan Bilangan Unik");

        System.out.println("Masukkkan Sejumlah Bilangan Unik : ");
        int inputBilangan = input.nextInt();

        // Validasi input bilangan unik
        if (inputBilangan < 1) {
            System.out.println("Jumlah bilangan harus lebih dari 0.");
        } else {
            int[] bilanganUnik = new int[inputBilangan];
            boolean isValid = true;

            for (int i = 0; i < inputBilangan; i++) {
                System.out.print("Masukkan bilangan ke-" + (i + 1) + ": ");
                int bilangan = input.nextInt();

                // Cek apakah bilangan sudah ada sebelumnya
                for (int j = 0; j < i; j++) {
                    if (bilanganUnik[j] == bilangan) {
                        System.out.println("Bilangan " + bilangan + " sudah ada, silakan masukkan bilangan unik lainnya.");
                        isValid = false;
                        i--; // Ulangi input untuk indeks yang sama
                        break;
                    }
                }

                if (isValid) {
                    bilanganUnik[i] = bilangan;
                } else {
                    isValid = true; // Reset untuk iterasi berikutnya
                }
            }

            System.out.println("Bilangan unik yang dimasukkan:");
            for (int bil : bilanganUnik) {
                System.out.print(bil + " ");
            }
        }
    }
}
