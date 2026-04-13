package com.example.addressbook.model;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import jakarta.validation.constraints.NotBlank;

@Entity
@Table(name = "address")
public class Address {

    @Id
    @NotBlank(message = "name is required")
    @Column(name = "name", nullable = false, length = 100)
    private String name;

    @NotBlank(message = "address_line1 is required")
    @Column(name = "address_line1", nullable = false, length = 255)
    private String addressLine1;

    @NotBlank(message = "city is required")
    @Column(name = "city", nullable = false, length = 100)
    private String city;

    @NotBlank(message = "state is required")
    @Column(name = "state", nullable = false, length = 100)
    private String state;

    @NotBlank(message = "country is required")
    @Column(name = "country", nullable = false, length = 100)
    private String country;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getAddressLine1() {
        return addressLine1;
    }

    public void setAddressLine1(String addressLine1) {
        this.addressLine1 = addressLine1;
    }

    public String getCity() {
        return city;
    }

    public void setCity(String city) {
        this.city = city;
    }

    public String getState() {
        return state;
    }

    public void setState(String state) {
        this.state = state;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }
}
