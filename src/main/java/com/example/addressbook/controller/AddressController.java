package com.example.addressbook.controller;

import com.example.addressbook.model.Address;
import com.example.addressbook.service.AddressService;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/addresses")
public class AddressController {

    private final AddressService addressService;

    public AddressController(AddressService addressService) {
        this.addressService = addressService;
    }

    @GetMapping
    public List<Address> getAllAddresses() {
        return addressService.getAll();
    }

    @GetMapping("/{name}")
    public Address getAddress(@PathVariable String name) {
        return addressService.getByName(name);
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Address createAddress(@Valid @RequestBody Address address) {
        return addressService.create(address);
    }

    @PutMapping("/{name}")
    public Address updateAddress(@PathVariable String name, @Valid @RequestBody Address address) {
        return addressService.update(name, address);
    }

    @DeleteMapping("/{name}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void deleteAddress(@PathVariable String name) {
        addressService.delete(name);
    }
}
