import java.util.Scanner;

public class konversiSuhuCelciusFahrenheit {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Konversi Suhu Celcius ke Fahrenheit");

        System.out.print("Masukkan suhu dalam Celcius: ");
        double suhuCelcius = input.nextDouble();

        // Validasi input
        if (suhuCelcius < -273.15) {
            System.out.println("Suhu tidak boleh kurang dari -273.15°C (nol mutlak).");
            return;
        }

        // Menghitung konversi suhu dari Celcius ke Fahrenheit
        double perhitunganSuhu = (9.0 / 5.0) * suhuCelcius + 32;
        System.out.println("Suhu dalam Fahrenheit adalah: " + perhitunganSuhu + "°F");
    }
}
