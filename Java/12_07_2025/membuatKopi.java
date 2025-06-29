import java.util.Scanner;

public class membuatKopi {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.println("Program Membuat Kopi");

        System.out.println("Dalam program ini, Hitungan menggunakan satuan sendok makan.");

        System.out.print("Masukan jumlah kopi yang anda: ");
        int jumlahKopi = input.nextInt();

        System.out.print("Masukan jumlah gula yang anda punya: ");
        int jumlahGula = input.nextInt();

        System.out.println("Anda mempunyai jumlah kopi: " + jumlahKopi + " sendok makan.");
        System.out.println("Anda mempunyai jumlah gula: " + jumlahGula + " sendok makan.");
        System.out.println("Gula dan kopi berhasil disimpan.");

        System.out.println("Apakah anda ingin membuat kopi? (ya/tidak)");
        do {
            String pilihan = input.next();

            if (pilihan.equalsIgnoreCase("ya")) {
                System.out.println("Anda memilih untuk membuat kopi.");
                System.out.println("Silakan masukkan jumlah kopi dan gula yang ingin Anda gunakan.");

                System.out.print("Masukkan jumlah kopi yang ingin dibuat (dalam sendok makan): ");
                int kopiYangInginDibuat = input.nextInt();

                System.out.print("Masukkan jumlah gula yang ingin ditambahkan (dalam sendok makan): ");
                int gulaYangInginDitambahkan = input.nextInt();

                if (kopiYangInginDibuat <= jumlahKopi && gulaYangInginDitambahkan <= jumlahGula) {
                    jumlahKopi = jumlahKopi - kopiYangInginDibuat;
                    jumlahGula = jumlahGula - gulaYangInginDitambahkan;
                    System.out.println("Kopi berhasil dibuat dengan " + kopiYangInginDibuat + " sendok makan kopi dan " + gulaYangInginDitambahkan + " sendok makan gula.");
                    System.out.println("Sisa kopi: " + jumlahKopi + " sendok makan.");
                    System.out.println("Sisa gula: " + jumlahGula + " sendok makan.");
                    System.out.println("Terima kasih telah menggunakan program ini!");
                } else {
                    System.out.println("Jumlah kopi atau gula tidak mencukupi untuk membuat kopi.");
                }

                break;
            } else if (pilihan.equalsIgnoreCase("tidak")) {
                System.out.println("Anda memilih untuk tidak membuat kopi.");
                System.out.println("Terima kasih telah menggunakan program ini!");
                break;
            } else {
                System.out.println("Pilihan tidak valid. Silakan masukkan 'ya' atau 'tidak'.");
            }
        } while (true);

    }
}
