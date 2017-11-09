package model

import "sync"

var vehiclePool sync.Pool

func init() {
	vehiclePool.New = func() interface{} {
		return new(Vehicle)
	}
}

/**
 * Класс, частично определяющий технику. Содержит уникальный идентификатор техники, а также все поля техники,
 * значения которых могут изменяться в процессе игры.
 */
type VehicleUpdate struct {
	/**
	 * Уникальный идентификатор объекта.
	 */
	Id int64
	/**
	 * X-координата центра объекта. Ось абсцисс направлена слева направо.
	 */
	X float64
	/**
	 * Y-координата центра объекта. Ось ординат направлена сверху вниз.
	 */
	Y float64
	/**
	 * Текущая прочность или {@code 0}, если техника была уничтожена либо ушла из зоны видимости.
	 */
	Durability int
	/**
	 * Количество тиков, оставшееся до следующей атаки.
	 * Для совершения атаки необходимо, чтобы это значение было равно нулю.
	 */
	RemainingAttackCooldownTicks int
	/**
	 * {@code true} в том и только том случае, если эта техника выделена.
	 */
	Selected bool
	/**
	 * Группы, в которые входит эта техника.
	 */
	Groups []int
}

/**
 * Класс, определяющий технику. Содержит также все свойства круглых объектов.
 */
type Vehicle struct {
	CircularUnit
	/**
	 * Идентификатор игрока, которому принадлежит техника.
	 */
	PlayerId int64
	/**
	 * Текущую прочность.
	 */
	Durability int
	/**
	 * Максимальную прочность.
	 */
	MaxDurability int
	/**
	 * Максимальное расстояние, на которое данная техника может переместиться за один игровой тик,
	 * без учёта типа местности и погоды. При перемещении по дуге учитывается длина дуги,
	 * а не кратчайшее расстояние между начальной и конечной точками.
	 */
	MaxSpeed float64
	/**
	 * Максимальное расстояние (от центра до центра),
	 * на котором данная техника обнаруживает другие объекты, без учёта типа местности и погоды.
	 */
	VisionRange float64
	/**
	 * Квадрат максимального расстояния (от центра до центра),
	 * на котором данная техника обнаруживает другие объекты, без учёта типа местности и погоды.
	 */
	SquaredVisionRange float64
	/**
	 * Максимальное расстояние (от центра до центра),
	 * на котором данная техника может атаковать наземные объекты.
	 */
	GroundAttackRange float64
	/**
	 * Квадрат максимального расстояния (от центра до центра),
	 * на котором данная техника может атаковать наземные объекты.
	 */
	SquaredGroundAttackRange float64
	/**
	 * Максимальное расстояние (от центра до центра),
	 * на котором данная техника может атаковать воздушные объекты.
	 */
	AerialAttackRange float64
	/**
	 * Квадрат максимального расстояния (от центра до центра),
	 * на котором данная техника может атаковать воздушные объекты.
	 */
	SquaredAerialAttackRange float64
	/**
	 * Урон одной атаки по наземному объекту.
	 */
	GroundDamage int
	/**
	 * Урон одной атаки по воздушному объекту.
	 */
	AerialDamage int
	/**
	 * Защиту от атак наземных юнитов.
	 */
	GroundDefence int
	/**
	 * Защиту от атак воздушых юнитов.
	 */
	AerialDefence int
	/**
	 * Минимально возможный интервал между двумя последовательными атаками данной техники.
	 */
	AttackCooldownTicks int
	/**
	 * Количество тиков, оставшееся до следующей атаки.
	 * Для совершения атаки необходимо, чтобы это значение было равно нулю.
	 */
	RemainingAttackCooldownTicks int
	/**
	 * Тип техники.
	 */
	VehicleType VehicleType
	/**
	 * {@code true} в том и только том случае, если эта техника воздушная.
	 */
	Aerial bool
	/**
	 * {@code true} в том и только том случае, если эта техника выделена.
	 */
	Selected bool
	/**
	 * Группы, в которые входит эта техника.
	 */
	Groups []int
}

func (v *Vehicle) Update() {

}