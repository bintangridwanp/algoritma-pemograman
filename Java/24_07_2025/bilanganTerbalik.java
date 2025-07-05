import java.util.Scanner;

public class bilanganTerbalik {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menampilkan Bilangan Terbalik");

        System.out.print("Masukkan bilangan bulat positif: ");
        int bilangan = input.nextInt();

        // Validasi input
        if (bilangan < 0) {
            System.out.println("Bilangan harus positif.");
            return;
        } else if (bilangan == 0) {
            System.out.println("Bilangan terbalik dari 0 adalah 0.");
            return;
        }

        int angka;
       while (bilangan > 0) {
           angka = bilangan % 10;
           System.out.print(angka);
           bilangan /= 10;
       }
    }
}
