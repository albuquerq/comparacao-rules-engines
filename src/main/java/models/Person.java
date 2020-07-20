package models;

public class Person {
    private String name;

    public Person(String name) {
        this.name = name;
    }

    public String getName() {
        System.out.println("Oi mundo!");
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
