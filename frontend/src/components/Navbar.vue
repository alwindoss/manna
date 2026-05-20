<template>
    <nav>
        <ul class="navbar">
            <li>
                <RouterLink to="/">Home</RouterLink>
            </li>
            <li>
                <RouterLink to="/settings">Settings</RouterLink>
            </li>
        </ul>
    </nav>
</template>

<style scoped>
.navbar {
    display: flex;
    gap: 3px;
    list-style-type: none;
}
</style>
