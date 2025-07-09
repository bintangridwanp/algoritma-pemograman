import java.util.Scanner;

public class cekPasswordKuat {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Cek Password Kuat");

        System.out.println("Masukkan password: ");
        String password = input.nextLine();

        // Validasi panjang password
        if (password.length() < 8) {
            System.out.println("Password harus memiliki minimal 8 karakter.");
        } else {
            // Validasi karakter khusus
            boolean hasUpperCase = false;
            boolean hasLowerCase = false;
            boolean hasDigit = false;
            boolean hasSpecialChar = false;

            for (char c : password.toCharArray()) {
                if (Character.isUpperCase(c)) {
                    hasUpperCase = true;
                } else if (Character.isLowerCase(c)) {
                    hasLowerCase = true;
                } else if (Character.isDigit(c)) {
                    hasDigit = true;
                } else if ("!@#$%^&*()-_=+[]{}|;:',.<>?/".indexOf(c) >= 0) {
                    hasSpecialChar = true;
                }
            }

            // Cek apakah password kuat
            if (hasUpperCase && hasLowerCase && hasDigit && hasSpecialChar) {
                System.out.println("Password Anda kuat.");
            } else {
                System.out.println("Password Anda lemah. Pastikan mengandung huruf besar, huruf kecil, angka, dan karakter khusus.");
            }
        }
    }
}
