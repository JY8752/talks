package org.example;

public class Item {
    private final String id;

    private Item(String id) {
        this.id = id;
    }

    public static Item createItem(String id) {
        if (id.startsWith("d-")) {
            return null;
        }
        return new Item(id);
    }
}
