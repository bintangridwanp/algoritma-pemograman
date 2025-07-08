//import java.util.Scanner;

public class menghitungJumlahUnik {
    public static void main(String[] args) {
        //Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Jumlah Unik dari Deret Angka");

        int[] angka = {10, 20, 10, 30, 20, 40, 50, 10};
        int jumlahUnik = 0;

        boolean cekUnik;

        for (int i = 0; i < angka.length; i++) {

            cekUnik = true;

            for (int j = 0; j < i; j++) {
                if (angka[j] == angka[i]) {
                        cekUnik = false;
                        break;
                }
            }

            if (cekUnik) {
                jumlahUnik++;
            }
        }

        System.out.println("Jumlah angka unik dalam deret adalah: " + jumlahUnik);
    }
}
