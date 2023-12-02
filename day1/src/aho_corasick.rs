/// The size of our alphabet,
/// all characters a-z plus digits 0-9
const ALPHABET_SIZE: u8 = 26 + 9;

/// An implementation of the Aho-Corasick algorithm for efficiently
/// searching for strings within the input text.
///
/// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
struct Automaton {
    states: Vec<State>,
    alphabet_size: usize,
}


impl Automaton {
    fn new(alphabet_size: usize) -> Self {
        Automaton {
            states: vec![State::new(alphabet_size)],
            alphabet_size,
        }
    }

    fn insert(mut self, v: &str, value: u8) -> Automaton {
        let ref mut states = self.states;

        self
    }

    fn search(&self, k: &str) -> Option<u8> {
        None
    }
}

struct State {
    /// Stores all transitions that can be made
    /// from this state, indexed by the alphabet
    /// offset of the character of the transition
    transitions: Vec<Option<usize>>,

    /// Stores the value of a final state,
    /// which is the value of the digit that was found
    value: Option<u8>,
}

impl State {
    fn new(alphabet_size: usize) -> Self {
        State {
            transitions: Vec::with_capacity(alphabet_size),
            value: None,
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::*;

    fn create_automaton() -> Automaton {
        Automaton::new(ALPHABET_SIZE as usize)
            .insert("one", 1)
            .insert("1", 1)
            .insert("two", 2)
            .insert("2", 2)
            .insert("three", 3)
            .insert("3", 3)
            .insert("four", 4)
            .insert("4", 4)
            .insert("five", 5)
            .insert("5", 5)
            .insert("six", 6)
            .insert("6", 6)
            .insert("seven", 7)
            .insert("7", 7)
            .insert("eight", 8)
            .insert("8", 8)
            .insert("nine", 9)
            .insert("9", 9)
    }

    #[test]
    fn test_simple() {
        let automaton = create_automaton();
        assert_eq!(Some(1), automaton.search("one"));
        assert_eq!(Some(1), automaton.search("1"));
        assert_eq!(Some(2), automaton.search("two"));
        assert_eq!(Some(2), automaton.search("2"));
        assert_eq!(Some(3), automaton.search("three"));
        assert_eq!(Some(3), automaton.search("3"));
        assert_eq!(Some(4), automaton.search("four"));
        assert_eq!(Some(4), automaton.search("4"));
        assert_eq!(Some(5), automaton.search("five"));
        assert_eq!(Some(5), automaton.search("5"));
        assert_eq!(Some(6), automaton.search("six"));
        assert_eq!(Some(6), automaton.search("6"));
        assert_eq!(Some(7), automaton.search("seven"));
        assert_eq!(Some(7), automaton.search("7"));
        assert_eq!(Some(8), automaton.search("eight"));
        assert_eq!(Some(8), automaton.search("8"));
        assert_eq!(Some(9), automaton.search("nine"));
        assert_eq!(Some(9), automaton.search("9"));
    }

    //noinspection SpellCheckingInspection
    #[test]
    fn test_extra() {
        let automaton = create_automaton();
        assert_eq!(Some(6), automaton.search("sixandfour"));
        assert_eq!(Some(5), automaton.search("five2"));
    }

    //noinspection SpellCheckingInspection
    #[test]
    fn test_combined() {
        let automaton = create_automaton();
        assert_eq!(Some(8), automaton.search("seveight"));
        assert_eq!(Some(1), automaton.search("fone"));
        assert_eq!(Some(3), automaton.search("nin3"));
    }

    //noinspection SpellCheckingInspection
    #[test]
    fn test_none() {
        let automaton = create_automaton();
        assert_eq!(None, automaton.search("thre"));
        assert_eq!(None, automaton.search("thre"));
    }
}
