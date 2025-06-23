import java.util.Scanner;

public class menentukanPositifNegatif {
    public static void main(String[] args) {
        Scanner inoput = new Scanner(System.in);
        System.out.println("Program Menentukan Bilangan Positif atau Negatif");

        System.out.print("Masukkan bilangan: ");
        int bilangan = inoput.nextInt();
        if (bilangan > 0) {
            System.out.println(bilangan + " adalah bilangan positif.");
        } else if (bilangan < 0) {
            System.out.println(bilangan + " adalah bilangan negatif.");
        } else {
            System.out.println("Bilangan yang dimasukkan adalah nol, bukan positif atau negatif.");
        }

    }
}
