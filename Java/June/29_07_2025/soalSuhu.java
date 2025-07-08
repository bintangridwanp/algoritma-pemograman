import java.util.Scanner;

public class soalSuhu {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Konversi Suhu");

        System.out.print("Masukkan suhu dalam Celcius : ");
        double suhuCelcius = input.nextDouble();

        double fahrenheit, kelvin, reamur;

        // validasi input suhu
        if (suhuCelcius < -273.15) {
            System.out.println("Suhu tidak valid. Suhu tidak boleh kurang dari -273.15 derajat Celcius.");
        } else {
            reamur = suhuCelcius * (4.0 / 5.0);
            fahrenheit = suhuCelcius * 9.0 / 5.0 + 32.0;
            kelvin = suhuCelcius + 273.15;

            System.out.println("Suhu dalam Reamur: " + reamur);
            System.out.println("Suhu dalam Fahrenheit: " + fahrenheit);
            System.out.println("Suhu dalam Kelvin: " + kelvin);

        }
    }
}
