/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

/** Infers the status of lasagna in oven from the timer
*
* @param {number} timer
* @returns {string} status of lasagna
*/
export function cookingStatus(timer) {
  if (timer == undefined) {
    return "You forgot to set the timer."
  }
  if (!timer) {
    return "Lasagna is done."
  } else {
    return "Not done, please wait."
  }
}

/** Estimates the preperation time
*
* @param {string[]} layers of lasagna to cook
* @param {number} preperation time per layer (minutes)
* @returns {string} total preperation time (minutes)
*/
export function preparationTime(layers, time) {
  return layers.length * (time ?? 2)
}

/**
 * Calculates the amount of noodles and sauce needed for a lasagna.
 *
 * @param {string[]} An array representing the lasagna layers.
 * @returns {{ noodles: number, sauce: number }} The required quantities of noodles (grams) and sauce (liters).
 */
export function quantities(layers) {
  const quantity = { noodles: 0, sauce: 0 };
  for (let i = 0; i < layers.length; i++) {
    if (layers[i] === "noodles") {
      quantity.noodles += 50;
    }
    if (layers[i] === "sauce") {
      quantity.sauce += 0.2;
    }
  }
  return quantity
}

/**
 * Adds the last index from my friends ingredients list to my ingredients list
 *
 * @param {number[]} friendsList
 * @param {number[]} myList
 */
export function addSecretIngredient(friendsList, myList) {
  myList.push(friendsList[friendsList.length-1])
}

/** Calculates the recipe for desired portions
 *
 * Assumes the original recipe serves 2 portions.
 *
 * @param {Object.<string, number>} Original ingredient names and values
 * @param {number} portions - The desired number of portions.
 * @returns {Object.<string, number>} A new recipe object with scaled ingredient quantities.
 */
export function scaleRecipe(recipe, portions) {
  const newRecipe = {}
  for (let key in recipe) {
    newRecipe[key] = recipe[key] * portions / 2
  }
  return newRecipe
}
