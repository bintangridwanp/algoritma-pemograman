import java.util.Scanner;

public class menghitungFaktorialBilangan {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Faktorial Bilangan");

        System.out.println("Masukkan bilangan bulat positif untuk menghitung faktorialnya: ");
        int bilangan = input.nextInt();

        // Validasi input
        if (bilangan < 0) {
            System.out.println("Faktorial tidak terdefinisi untuk bilangan negatif.");
            return;
        }

        // Menghitung faktorial
        int hasil = 1;
        for (int i = 1; i <= bilangan; i++) {

            if (i == 1) {
                System.out.print(i);
            } else {
                hasil = hasil * i;
                System.out.print(" x " + i);
            }
        }
        System.out.print(" = " + hasil);
    }
}
