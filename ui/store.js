import Vuex from "vuex";
import axios from "axios";

export const store = new Vuex.Store({
	strict: process.env.NODE_ENV !== 'production',
	state: {
		loading: false,
		user: null,
		users: [],
		repositories: [],
	},
	getters: {
		getUser: (state) => (username) => {
			const users = state.users.filter((user) => {
				return user.username === username;
			});
			if (users.length > 0) {
				return users[0];
			}
			return {};
		}
	},
	mutations: {
		loading(state, isLoading) {
			state.loading = isLoading;
		},
		setUser(state, user) {
			state.user = user;
		},
		addUsers(state, users) {
			state.users = users;
		},
		addUser(state, newUser) {
			for (let i = 0; i < state.users.length; i++) {
				if (state.users[i].username === newUser.username) {
					state.users[i] = newUser;
					return
				}
			}
			state.users.push(newUser);
		},
		updateUser(state, updatedUser) {
			for (let i = 0; i < state.users.length; i++) {
				if (state.users[i].id === updatedUser.id) {
					state.users[i] = updatedUser;

					// If the current user was updated, update it in the store too
					if (state.user.id === updatedUser.id) {
						state.user = updatedUser;
					}

					return
				}
			}
		},
		setRepositories(state, repositories) {
			state.repositories = repositories;
		},
	},
	actions: {
		fetchAuthenticatedUser(ctx) {
			return new Promise((resolve, reject) => {
				axios.get(`${window.config.api}/user`)
					.then((res) => {
						ctx.commit('setUser', res.data);
						resolve(res.data);
					})
					.catch((err) => {
						reject(err);
					})
			})
		},
		fetchUserRepositories(ctx, username) {
			return new Promise((resolve, reject) => {
				axios.get(`${window.config.api}/users/${username}/repositories`)
					.then((res) => {
						ctx.commit('setRepositories', res.data);
						resolve(res.data);
					})
					.catch((err) => {
						reject(err);
					})
			})
		},
		authenticateUser(ctx, payload) {
			return new Promise((resolve, reject) => {
				axios.post(`${window.config.api}/authorize`, payload)
					.then((res) => {
						ctx.commit('setUser', res.data);
						resolve(res.data);
					})
					.catch((err) => {
						reject(err);
					})
			})
		},
		fetchUsers(ctx) {
			axios.get(`${window.config.api}/users`)
				.then((res) => {
					ctx.commit('addUsers', res.data);
				})
				.catch((err) => {
					alert(err);
				})
		},
		fetchUser(ctx, username) {
			axios.get(`${window.config.api}/users/${username}`)
				.then((res) => {
					ctx.commit('addUser', res.data);
				})
				.catch((err) => {
					alert(err);
				})
		},
		updateUser(ctx, user) {
			return new Promise((resolve, reject) => {
				axios.put(`${window.config.api}/users/${user.username}`, user)
					.then((res) => {
						ctx.commit('updateUser', res.data);
						resolve(res.data);
					})
					.catch((err) => {
						reject(err);
					})
			})
		},
		deleteUser(ctx, username) {
			axios.delete(`${window.config.api}/users/${username}`)
				.then((res) => {
					ctx.dispatch('fetchUsers');
				})
				.catch((err) => {
					alert(err);
				})
		}
	},
});
