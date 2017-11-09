package model

type PlayerContext struct {
	*Player
	*World
}

/**
 * Содержит данные о текущем состоянии игрока.
 */
type Player struct {
	/**
	 * Уникальный идентификатор игрока.
	 */
	Id int64
	/**
	 * {@code true} в том и только в том случае, если этот игрок ваш.
	 */
	Me bool
	/**
	 * Специальный флаг --- показатель того, что стратегия игрока <<упала>>.
	 * Более подробную информацию можно найти в документации к игре.
	 */
	StrategyCrashed bool
	/**
	 * Количество баллов, набранное игроком.
	 */
	Score int
	/**
	 * Количество тиков, оставшееся до любого следующего действия.
	 * Если значение равно {@code 0}, игрок может совершить действие в данный тик.
	 */
	RemainingActionCooldownTicks int
}