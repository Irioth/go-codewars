package codewars

/**
 * Стратегия --- интерфейс, содержащий описание методов искусственного интеллекта армии.
 * Каждая пользовательская стратегия должна реализовывать этот интерфейс.
 * Может отсутствовать в некоторых языковых пакетах, если язык не поддерживает интерфейсы.
 */
type Strategy interface {
	/**
		 * Основной метод стратегии, осуществляющий управление армией. Вызывается каждый тик.
		 *
	     * me    Информация о вашем игроке.
	     * world Текущее состояние мира.
	     * game  Различные игровые константы.
	     * move  Результатом работы метода является изменение полей данного объекта.
	*/
	Move(*Player, *World, *Game, *Move)
}
