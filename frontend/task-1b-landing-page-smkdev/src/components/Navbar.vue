<script setup>

</script>

<template>
    <nav class="navbar navbar-expand-lg bg-body-tertiary fixed-top">
        <div class="container">
            <img src="./assets/images/smkdev-logo.png" style="width: 100px;"/>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
                aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse d-lg-flex justify-content-lg-end" id="navbarNav">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <a class="nav-link" aria-current="page" href="#">Home</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Learn</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Community</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Blog</a>
                    </li>
                    <li class="nav-item d-flex justify-content-center justify-content-lg-start">
                        <button class="nav-link text-white fw-semibold px-3 ms-lg-3" style="border: 1px solid; border-radius: 10px; background: rgb(23,205,72);
background: linear-gradient(144deg, rgba(23,205,72,1) 45%, rgba(255,193,107,1) 88%);">Dashboard</button>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>