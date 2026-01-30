<template>
  <v-navigation-drawer expand-on-hover permanent rail>
    <v-list>
      <v-list-item
        prepend-avatar="https://randomuser.me/api/portraits/women/85.jpg"
        subtitle="sandra_a88@gmailcom"
        title="Sandra Adams"
        :to="'/'"
      />
    </v-list>

    <v-divider />

    <v-list density="compact" nav>
      <v-list-item
        prepend-icon="mdi-folder"
        title="My Files"
        :to="'/about'"
        value="myfiles"
      />
      <v-list-item
        prepend-icon="mdi-account-multiple"
        title="Shared with me"
        value="shared"
      />
      <v-list-item prepend-icon="mdi-star" title="Starred" value="starred" />
    </v-list>

    <v-divider />

    <v-list density="compact" nav>
      <v-list-item
        prepend-icon="mdi-format-list-checks"
        title="Task List"
        :to="'/tasklist'"
        value="tasklist"
      />
      <v-list-item
        prepend-icon="mdi-account-group"
        title="Participants"
        :to="'/participants'"
        value="participants"
      />
    </v-list>
    <v-list-group prepend-icon="mdi-cog" title="Settings" value="settings">
      <v-list-item title="settings" :to="'/settings'" value="settings" />
    </v-list-group>
  </v-navigation-drawer>
</template>
