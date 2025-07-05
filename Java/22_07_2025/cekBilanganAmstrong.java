import java.util.Scanner;

public class cekBilanganAmstrong {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Cek Bilangan Armstrong");
        System.out.print("Masukkan bilangan bulat positif: ");
        int masukkanBilangan = input.nextInt();

        if (masukkanBilangan == 0) {
            System.out.println("Bilangan Armstrong tidak terdefinisi untuk 0.");
            return;
        } else if (masukkanBilangan < 0) {
            System.out.println("Bilangan harus positif.");
            return;
        }

        int jumlahDigit = hitungJumlahDigit(masukkanBilangan);
        int jumlahPangkat = hitungJumlahPangkat(masukkanBilangan, jumlahDigit);

        if (jumlahPangkat == masukkanBilangan) {
            System.out.println(masukkanBilangan + " adalah bilangan Armstrong.");
        } else {
            System.out.println(masukkanBilangan + " bukan bilangan Armstrong.");
        }
    }

    // Menghitung jumlah digit pada bilangan
    static int hitungJumlahDigit(int bilangan) {
        int jumlah = 0;
        int temp = bilangan;
        while (temp > 0) {
            jumlah++;
            temp /= 10;
        }
        return jumlah;
    }

    // Menghitung jumlah pangkat dari setiap digit
    static int hitungJumlahPangkat(int bilangan, int jumlahDigit) {
        int jumlahPangkat = 0;
        int temp = bilangan;
        while (temp > 0) {
            int digit = temp % 10;
            int pangkat = 1;
            for (int i = 0; i < jumlahDigit; i++) {
                pangkat *= digit;
            }
            jumlahPangkat += pangkat;
            temp /= 10;
        }
        return jumlahPangkat;
    }
}