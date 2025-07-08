import java.util.Scanner;

public class kelulusan {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Kelulusan Mahasiswa");

        System.out.print("Masukan nilai mahasiswa : ");
        int nilai = input.nextInt();

        System.out.print("Konfirmasi mahassiswa mengerjakan ujian (true/false) : ");
        boolean tugas = input.nextBoolean();

        if (nilai > 55 && tugas) {
            System.out.println("Selamat, Anda lulus ujian!, dengan nilai: " + nilai + " dan tugas: " + tugas);
        } else if (nilai > 90) {
            System.out.println("Selamat, Anda lulus ujian dengan nilai sangat baik: " + nilai );
        } else {
            System.out.println("Anda tidak lulus ujian, dengan nilai: " + nilai + " dan tugas: " + tugas);
        }
        input.close();
    }
}
