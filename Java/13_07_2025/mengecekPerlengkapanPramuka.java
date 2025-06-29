import java.util.Scanner;

public class mengecekPerlengkapanPramuka {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.print("Masukkan jumlah anggota tim: ");
        int N = input.nextInt();
        input.nextLine(); // consume newline

        boolean lengkap = true;

        for (int i = 1; i <= N; i++) {
            System.out.println("Cek kelengkapan anggota ke-" + i + ":");
            System.out.print("Topi (y/n): ");
            String topi = input.nextLine();
            System.out.print("Ikat pinggang (y/n): ");
            String ikatPinggang = input.nextLine();
            System.out.print("Alat tulis (y/n): ");
            String alatTulis = input.nextLine();
            System.out.print("Tali (y/n): ");
            String tali = input.nextLine();
            System.out.print("Pisau pramuka (y/n): ");
            String pisau = input.nextLine();

            if (!topi.equalsIgnoreCase("y") || !ikatPinggang.equalsIgnoreCase("y") ||
                    !alatTulis.equalsIgnoreCase("y") || !tali.equalsIgnoreCase("y") ||
                    !pisau.equalsIgnoreCase("y")) {
                lengkap = false;
            }
        }

        if (lengkap) {
            System.out.println("Persiapan tim lengkap");
        } else {
            System.out.println("Persiapan tim tidak lengkap");
        }
    }
}