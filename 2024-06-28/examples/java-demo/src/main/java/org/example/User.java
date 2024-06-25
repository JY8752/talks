package org.example;

public class User {

    private final String name;
    private final int age;

    // コンストラクタ
    private User(String name, int age) {
        this.name  = name;
        this.age = age;
    }

    // ファクトリ関数
    public static User createUser(String name, int age) {
        return new User(name, age);
    }

    // getter
    public String getName() {
        return name;
    }

    public int getAge() {
        return age;
    }
}
