import java.util.Scanner;

public class mendapatkanBola {

    public static void tampilkanWarnaBola(String bolaPertama, String bolaKedua, String bolaKetiga) {
        System.out.println("Bola pertama: " + bolaPertama);
        System.out.println("Bola kedua: " + bolaKedua);
        System.out.println("Bola ketiga: " + bolaKetiga);
    }

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Mendapatkan Bola");

        System.out.print("Masukkan warna bola pertama: ");
        String bolaPertama = input.nextLine();

        System.out.print("Masukkan warna bola kedua: ");
        String bolaKedua = input.nextLine();

        System.out.print("Masukkan warna bola ketiga: ");
        String bolaKetiga = input.nextLine();

        System.out.println("Warna bola berhasil di simpan.");

        String pilihan;
        do {
            System.out.println("Apakah anda ingin memulai program untuk merubah warna bola? (y/n)");
            pilihan = input.nextLine();

            if (pilihan.equals("y") || pilihan.equals("iya")) {
                System.out.println("Anda memilih untuk mengubah tempat bola.");
                tampilkanWarnaBola(bolaPertama, bolaKedua, bolaKetiga);

                System.out.println("Masukkan nomer bola yang ingin anda rubah warnanya: ");
                int nomerBola = input.nextInt();
                input.nextLine(); // consume newline

                if (nomerBola == 1) {
                    System.out.println("Masukkan warna bola pertama yang baru: ");
                    bolaPertama = input.nextLine();
                } else if (nomerBola == 2) {
                    System.out.println("Masukkan warna bola kedua yang baru: ");
                    bolaKedua = input.nextLine();
                } else if (nomerBola == 3) {
                    System.out.println("Masukkan warna bola ketiga yang baru: ");
                    bolaKetiga = input.nextLine();
                } else {
                    System.out.println("Pilihan tidak valid. Silakan masukkan 1, 2, atau 3.");
                }
                tampilkanWarnaBola(bolaPertama, bolaKedua, bolaKetiga);
            } else if (pilihan.equals("n") || pilihan.equals("tidak")) {
                System.out.println("Anda memilih untuk tidak mengubah warna bola.");
                tampilkanWarnaBola(bolaPertama, bolaKedua, bolaKetiga);
                System.out.println("Terima kasih telah menggunakan program ini!");
                break;
            } else {
                System.out.println("Pilihan tidak valid. Silakan masukkan 'y' untuk ya atau 'n' untuk tidak.");
            }
        } while (true);
    }
}