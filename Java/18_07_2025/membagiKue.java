import java.util.Scanner;

public class membagiKue {
    public static void main(String[] args) {
        Scanner intput = new Scanner(System.in);

        System.out.println("Program Membagi Kue");

        System.out.print("Masukkan jumlah kue yang ingin dibagi: ");
        int jumlahKue = intput.nextInt();

        System.out.print("Masukkan jumlah orang yang akan menerima kue: ");
        int jumlahOrang = intput.nextInt();

        // Memeriksa apakah jumlah kue dan orang valid
        if (jumlahKue < 0 || jumlahOrang <= 0) {
            System.out.println("Jumlah kue harus positif dan jumlah orang harus lebih dari nol.");
            return;
        }

        // Menghitung jumlah kue per orang
        int kuePerOrang = jumlahKue / jumlahOrang;
        // Menghitung sisa kue
        int sisaKue = jumlahKue % jumlahOrang;

        // Menampilkan hasil pembagian kue
        System.out.println("Setiap orang akan menerima " + kuePerOrang + " kue.");
        if (sisaKue > 0) {
            System.out.println("Sisa kue yang tidak terbagi adalah " + sisaKue + " kue.");
        } else {
            System.out.println("Tidak ada sisa kue yang tidak terbagi.");
        }


    }
}
