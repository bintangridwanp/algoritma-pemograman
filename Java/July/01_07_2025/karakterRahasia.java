import java.util.Scanner;

public class karakterRahasia {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Karakter Rahasia");

        System.out.print("Masukkan x1: ");
        int x1 = input.nextInt();

        System.out.print("Masukkan x2: ");
        int x2 = input.nextInt();

        System.out.print("Masukkan x3: ");
        int x3 = input.nextInt();

        System.out.print("Masukkan x4: ");
        int x4 = input.nextInt();

        System.out.print("Masukkan x5: ");
        int x5 = input.nextInt();

        // Hitung nilai karakter rahasia
        char y1 = (char)((((x1 + x2) % 4096) >> 6) + 64);
        char y2 = (char)((((x2 + x3) % 4096) >> 6) + 64);
        char y3 = (char)((((x3 + x4) % 4096) >> 6) + 64);
        char y4 = (char)((((x4 + x5) % 4096) >> 6) + 64);

        // Tampilkan hasil karakter rahasia
        System.out.printf("Karakter rahasia: %c%c%c%c\n", (char) y1, (char) y2, (char) y3, (char) y4);

    }
}
